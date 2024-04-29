const express = require("express")

const pool = require("../database/db")

const router = express.Router()

router.post("/", async (req, res) => {
  try {
    const body = req.body
    if (!body.email, !body.username, !body.password) {
      res.status(400).json(err)
    }
    let id = "asdfasdfasdfasdfasdf" 
    const new_user = await pool.query("INSERT INTO users (id, username, email, password, owner_of) VALUES ($1, $2, $3, $4, $5);", [id, body.name, body.username, body.password, body.owner_of])
    console.log(new_user)
    res.status(201).json({"username": new_user.username})
  } catch(err) {
    res.status(500).json(err)
  }
})

router.get("/", async (req, res) => {
  try {
    // const user_by_name = await User.findAll()
    res.status(200).json(user_by_name)
  } catch(err) {
    res.status(500).json(err)
  }
})

router.get("/name", async (req, res) => {
  try {
    const name = req.body
    console.log(name)
    // const user_by_name = await User.findOne(name)
    res.status(200).json(name)
  } catch(err) {
    res.status(500).json(err)
  }
})

function routerApi(app) {
  app.use("/api/users", router)
}

module.exports = routerApi
