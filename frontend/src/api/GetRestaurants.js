import axios from 'axios'

export default function GetRestaurants() {
  return axios.get("http://localhost:8000/api/restaurants")
  .then((res) => console.log(res))
  .catch((err) => console.log(err))
}
