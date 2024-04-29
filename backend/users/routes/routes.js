const express = require("express")

const pool = require("../database/db")

const router = express.Router()

router.post("/", async (req, res) => {
  try {
    const body = req.body
    if (!body.email, !body.username, !body.password) {
      res.status(400).json(err)
    }
    let id = "asdfasdfasdfasdf" 
    const new_user = await pool.query("INSERT INTO users (id, username, email, password, owner_of) VALUES ($1, $2, $3, $4, $5);", [id, body.username, body.email, body.password, body.owner_of])
    res.status(201).json("User inserted successfully")
  } catch(err) {
    res.status(500).json(err)
  }
})

router.get("/", async (req, res) => {
  try {
    const users = await pool.query("SELECT username, email, owner_of FROM users;")
    res.status(200).json(users.rows)
  } catch(err) {
    res.status(500).json(err)
  }
})

router.get("/username", async (req, res) => {
  try {
    const username = req.body.username
    const user_by_username = await pool.query("SELECT email, username, owner_of FROM users WHERE username = $1;", [username])
    res.status(200).json(user_by_username.rows[0])
  } catch(err) {
    res.status(500).json(err)
  }
})

router.patch("/", async (req, res) => {
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
