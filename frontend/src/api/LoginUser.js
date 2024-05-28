import axios from 'axios'

export default function LoginUser(data) {
  return axios({
    url: 'http://localhost:8080/api/users/login',
    method: 'post',
    data: {
      email: data.email,
      password: data.password
    }
  })
  .then((res) => res.data)
  .catch((err) => console.log(err))
}