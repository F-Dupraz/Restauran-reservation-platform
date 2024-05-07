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
    const new_user = await pool.query("INSERT INTO users (id, username, email, password, owner_of) VALUES ($1, $2, $3, $4, $5);", [id, body.username, body.email, hashed_password, body.owner_of])
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

router.patch("/", authenticate({ throwOnError: true }), async (req, res) => {
  try {
    const body = req.body
    if (!body.username, !body.owner_of) {
      res.status(400).json(err)
    }
    const updated_user = await pool.query("UPDATE users SET owner_of=$1 WHERE username=$2;", [body.owner_of, body.username])
    res.status(200).json("User updated successfully")
  } catch (err) {
    res.status(500).json(err)
  }
})

function routerApi(app) {
  app.use("/api/users", router)
}

module.exports = routerApi
