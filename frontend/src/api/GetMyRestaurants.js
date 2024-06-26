import axios from 'axios'

export default function GetRestaurantById(token) {
  return axios({
    method: "get",
    url: "http://localhost:8000/api/my-restaurants",
    headers: {
      "Authorization": `${token}`,
      "Content-Type": "application/json"
    }
  })
  .then((res) => res.data.my_restaurants)
  .catch((err) => console.log(err))
}