import axios from 'axios'

export default function GetRestaurantById(idParam, token) {
  return axios({
    method: "get",
    url: `http://localhost:8020/api/reservations/${idParam}`,
    headers: {
      "Authorization": `${token}`,
      "Content-Type": "application/json"
    }
  })
  .then((res) => res.data)
  .catch((err) => console.log(err))
}