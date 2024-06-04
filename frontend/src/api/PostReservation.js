import axios from 'axios'

export default function PostReservation(data, token) {
  return axios({
    url: 'http://localhost:8020/api/reservations',
    method: 'post',
    data: {
      restaurant_id: data.restaurant_id,
      from: data.from,
      to: data.to,
      day: data.date,
      num_guests: data.num_guests,
    },
    headers: {
      "Authorization": `${token}`,
      "Content-Type": "application/json"
    }
  })
  .then((res) => res.data)
  .catch((err) => console.log(err))
}