import axios from 'axios'

export default function DeleRestaurant(res_id, token) {
  return axios({
    url: 'http://localhost:8020/api/reservations',
    method: 'delete',
    data: {
      id: res_id
    },
    headers: {
      "Authorization": `${token}`,
      "Content-Type": "application/json"
    }
  })
  .then((res) => res.data)
  .catch((err) => err.response.data)
}