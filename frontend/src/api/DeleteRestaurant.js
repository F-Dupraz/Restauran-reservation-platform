import axios from 'axios'

export default function DeleRestaurant(data, token) {
  return axios({
    url: 'http://localhost:8000/api/restaurants',
    method: 'delete',
    data: {
      id: data.id
    },
    headers: {
      "Authorization": `${token}`,
      "Content-Type": "application/json"
    }
  })
  .then((res) => res.data)
  .catch((err) => err.response.data)
}