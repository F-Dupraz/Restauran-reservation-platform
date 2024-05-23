import axios from 'axios'

export default function GetRestaurants(offsetParam) {
  return axios({
    method: "get",
    url: `http://localhost:8000/api/restaurants?offset=${offsetParam}`
  })
  .then((res) => res.data.restaurants)
  .catch((err) => console.log(err))
}
