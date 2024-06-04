import axios from 'axios'

export default function PostRestaurant(data, token) {
  return axios({
    url: 'http://localhost:8000/api/restaurants',
    method: 'post',
    data: {
      name: data.name,
      city: data.city,
      address: data.address,
      description: data.description,
      days_open: data.days_open,
      from: data.from,
      to: data.to,
      capacity: data.capacity,
      specialties: data.specialties,
    },
    headers: {
      "Authorization": `${token}`,
      "Content-Type": "application/json"
    }
  })
  .then((res) => res.data)
  .catch((err) => console.log(err))
}