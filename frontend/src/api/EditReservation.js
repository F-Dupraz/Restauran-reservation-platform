import axios from 'axios'

export default function EditReservation(data, token) {
  return axios({
    url: 'http://localhost:8020/api/reservations',
    method: 'patch',
    data: {
      id: data.id,
      day: data.day,
      from: data.from,
      to: data.to,
      num_guests: data.num_guests,
    },
    headers: {
      "Authorization": `${token}`,
      "Content-Type": "application/json"
    }
  })
  .then((res) => res.data)
  .catch((err) => err.response.data)
}