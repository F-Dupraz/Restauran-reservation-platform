import axios from 'axios'

export default function GetRestaurantById(usernameParam) {
  return axios({
    method: "get",
    url: `http://localhost:8080/api/users/${usernameParam}`
  })
  .then((res) => res.data)
  .catch((err) => console.log(err))
}