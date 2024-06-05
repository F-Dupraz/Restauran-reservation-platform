import axios from 'axios'

export default function GetRestaurantById(token) {
  return axios({
    method: "get",
    url: "http://localhost:8020/api/my-reservations",
    headers: {
      "Authorization": `${token}`,
      "Content-Type": "application/json"
    }
  })
  .then((res) => res.data.MyReservations)
  .catch((err) => console.log(err))
}