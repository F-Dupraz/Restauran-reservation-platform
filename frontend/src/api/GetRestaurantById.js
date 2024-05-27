import axios from 'axios'

export default function GetRestaurantById(idParam) {
  return axios({
    method: "get",
    url: `http://localhost:8000/api/restaurants/${idParam}`
  })
  .then((res) => res.data.restaurant)
  .catch((err) => console.log(err))
}