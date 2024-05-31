const { jwtVerify } = require("jose")
const { TextEncoder } = require("util")

const pool = require("../database/db")

const verifyToken = async (req) => {
  console.log(req)
  const { authorization } = req.headers
  const token = (authorization || '').replace('Bearer ', '')

  try {
    const verified = await jwtVerify(
      token,
      new TextEncoder().encode(process.env.JWT_TOKEN)
    )

    console.log(token)

    return verified.payload
  } catch (e) { 
    throw new Error('Invalid token')
  }
}

const defaultOptions = {
  throwOnError: true,
}

const authenticate = (options) => async (req, res, next) => {
  const _options = { ...defaultOptions, ...options }

  console.log(_options)

  try {
    const payload = await verifyToken(req)

    console.log(payload)

    const username = payload.username
    const user = await pool.query("SELECT id, username FROM users WHERE username = $1", [username])

    if (!user) {
      throw new Error('Invalid token')
    }

    req.user = user

    console.log(req)

    next()
  } catch (e) {
    if (e && e.message) {
      console.error(e.message)
    }

    if (_options.throwOnError) {
      return res.status(401).json({ error: e.message })
    }

    next()
  }
}

module.exports = authenticate
