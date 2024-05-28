import axios from 'axios'

export default function PostUser(data) {
  return axios({
    url: 'http://localhost:8080/api/users',
    method: 'post',
    data: {
      email: data.email,
      username: data.username,
      password: data.password
    }
  })
  .then((res) => console.log(res))
  .catch((err) => console.log(err))
}