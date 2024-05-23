const express = require("express")
const { v4: uuidv4 } = require('uuid')
const bcrypt = require("bcrypt")
const { SignJWT } = require("jose")
const { TextEncoder } = require("util")

const pool = require("../database/db")
const hashPassword = require("../helpers/auth")
const authenticate = require("../middlewares/auth")

const router = express.Router()

router.post("/", async (req, res) => {
  try {
    const body = req.body
    if (!body.email || !body.username || !body.password) {
      res.status(400).json(err)
    }
    let id = uuidv4()
    let hashed_password = await hashPassword(body.password)
    const new_user = await pool.query("INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4);", [id, body.username, body.email, hashed_password])
    res.status(201).json("User inserted successfully")
  } catch(err) {
    res.status(500).json(err)
  }
})

router.post("/login", async (req, res) => {
  try {
    const { email, password } = req.body
    if (!email || !password) {
      return res.status(400).json({
        error: "Username or password is not specified",
      })
    }
    const user = await pool.query("SELECT id, email, username, password FROM users WHERE email = $1;", [email])
    if (!user) {
      return res.status(400).json({
        error: "Username is incorrect",
      })
    }

    const is_pass_valid = await bcrypt.compare(password, user.rows[0].password)
    if (!is_pass_valid) {
      return res.status(401).json({
        error: "password is incorrect",
      })
    }

    const token = await new SignJWT({ id: user.rows[0].id, username: user.rows[0].username })
        .setProtectedHeader({ alg: "HS256" })
        .setIssuedAt()
        .setExpirationTime("2days")
        .sign(new TextEncoder().encode(process.env.JWT_TOKEN))

    return res.status(200).json({ username: user.rows[0].username, token: token, })
  } catch (err) {
    return res.status(500).json({
      "message": "An unexpected error happened. Please try again later",
      "error": err,
    })
  }
})

router.get("/", async (req, res) => {
  try {
    const users = await pool.query("SELECT id, username, email, owner_of FROM users;")
    res.status(200).json(users.rows)
  } catch(err) {
    res.status(500).json(err)
  }
})

router.get("/username", async (req, res) => {
  try {
    const username = req.body.username
    if (!username) {
      return res.status(400).json({
        error: "Username not specified",
      })
    }
    const user_by_username = await pool.query("SELECT id, email, username, owner_of FROM users WHERE username = $1;", [username])
    res.status(200).json(user_by_username.rows[0])
  } catch(err) {
    res.status(500).json(err)
  }
})

router.get("/reservations", authenticate({ throwOnError: true }), async (req, res) => {
  try {
    const user = req.user.rows[0].id
    if (!user) {
      res.status(400).json("Bad request. Maybe you forgot something.")
    }
    const my_reservations = await pool.query("SELECT day, h_from, h_to, num_guests FROM reservations WHERE user_id=$1", [user])
    res.status(200).json(my_reservations.rows)
  } catch(err) {
    res.status(500).json(err)
  }
})

router.get("/restaurants", authenticate({ throwOnError: true }), async (req, res) => {
  try {
    const user = req.user.rows[0].id
    if (!user) {
      res.status(400).json("Bad request. Maybe you forgot something.")
    }
    const my_restaurants = await pool.query("SELECT name, city, address, description, days_open, working_hours, capacity, specialties FROM restaurants WHERE owner=$1", [user])

    let parsed_restaurants = my_restaurants.rows.map(row => {
      let wh_string = row.working_hours.replace(/{/g, "")
        .replace(/}/g, "")
        .replace(/\[/g, "")
        .replace(/]/g, "")
        .replace(/\\/g, "")
        .replace(/"/g, "")
      let working_hours = wh_string.split(",")

      return { name: row.name, city: row.city, address: row.address, description: row.description, days_open: row.days_open, working_hours: working_hours, capacity: row.capacity, specialties: row.specialties }
    })

    res.status(200).json(parsed_restaurants)
  } catch(err) {
    console.log(err)
    res.status(500).json(err)
  }
})

function routerApi(app) {
  app.use("/api/users", router)
}

module.exports = routerApi
