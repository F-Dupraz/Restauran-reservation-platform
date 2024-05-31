import './UserPage.css'

import { useState, useEffect } from 'react'

import { Link } from 'react-router-dom'

import Header from '../../containers/Header/Header'
import Footer from '../../containers/Footer/Footer'

import GetUser from '../../api/GetUser'
import GetMyRestaurants from '../../api/GetMyRestaurants'
import GetMyReservations from '../../api/GetMyReservations'

import capitalizeFirstLetterOfEachWord from '../../helpers/capitalizeFunction.js'

export default function UserPage() {
  const getUsername = () => {
    return localStorage.getItem('my_username')
  }

  const getMyToken = () => {
    return localStorage.getItem('my_token')
  }

  const [user, setUser] = useState([])
  const [myRestaurants, setMyRestaurants] = useState([])
  const [myReservations, setMyReservations] = useState([])

  useEffect(() => {
    const fetchUser = async () => {
      const username = getUsername()
      const my_token = getMyToken()
      const userResponse = await GetUser(username)
      const myRestaurantsResponse = await GetMyRestaurants(my_token)
      const myReservationsResponse = await GetMyReservations(my_token)
      setUser(userResponse)
      setMyRestaurants(myRestaurantsResponse)
      setMyReservations(myReservationsResponse)
    }

    fetchUser()

    console.log(myRestaurants)
    console.log(myReservations)
  }, [])

  return (
    <>
      <Header />

      <main>
        <section className='userpage-section'>
          <div className='userpage-info'>
            <h2>{user.username}</h2>
            <p>{user.email}</p>
            <p>{user.id}</p>
          </div>
          <div className='userpage-restaurants'>
            <h3>Mis Restaurantes</h3>
            <div className='userpage-restaurant-container'>
              {
                myRestaurants.map((restaurant, index) => {
                  return (
                    <div className='userpage-restaurant' key={index}>
                      <Link to="/restaurants/:id">
                        <h4>{capitalizeFirstLetterOfEachWord(restaurant.name)}</h4>
                        <p><i>{restaurant.id}</i></p>
                      </Link>
                      <Link to="/edit-restaurants/:id" className='editrestaurant-link'>
                        Editar
                      </Link>
                    </div>
                  )
                }) 
              }
            </div>
            <p className='userpage-restaurants-p'>
              <Link to="/new-restaurant">Añadir Restaurante</Link>
            </p>
          </div>
          <div className='userpage-reservations'>
            <h3>Mis Reservas</h3>
            {
              myReservations.map((reservation, index) => {
                return (
                  <div className='userpage-reservation' key={index}>
                    <Link to="/reservations/:id">
                      <p>{capitalizeFirstLetterOfEachWord(reservation.restaurant_name)}</p>
                      <p>{reservation.day.slice(0,10)}</p>
                      <p><i>{reservation.id}</i></p>
                    </Link>
                  </div>
                )
              })
            }
          </div>
        </section>
      </main>

      <Footer />
    </>
  )
}
