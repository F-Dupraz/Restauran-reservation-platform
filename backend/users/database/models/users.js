const { DataTypes, Sequelize } = require('sequelize')
const sequelize = new Sequelize('sqlite::memory:')

const User = sequelize.define(
  "Users",
  {
    id: {
      allowNull: false,
      primaryKey: true,
      type: DataTypes.STRING
    }, 
    email: {
      allowNull: false,
      type: DataTypes.STRING,
      unique: true,
    },
    password: {
      allowNull: false,
      type: DataTypes.STRING
    },
    created_at: {
      allowNull: false,
      type: DataTypes.DATE,
      field: 'create_at',
      defaultValue: Sequelize.NOW
    }
  }
)

module.exports = User
