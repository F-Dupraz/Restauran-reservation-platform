const { genSalt, hash } = require("bcrypt")

const SALT_ROUNDS = 10

const hashPassword = async (password) => {
  const salt = await genSalt(SALT_ROUNDS)
  return hash(password, salt)
};

module.exports = hashPassword
