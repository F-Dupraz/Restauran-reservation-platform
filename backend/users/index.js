const express = require("express")
const cors = require("cors")
require("dotenv").config()

const routerApi = require("./routes/routes")

const app = express()
const PORT = process.env.USERS_PORT

app.use(express.json())
app.use(cors())

routerApi(app)

app.listen(PORT, () => {
  console.log("App running on port", PORT)
})
