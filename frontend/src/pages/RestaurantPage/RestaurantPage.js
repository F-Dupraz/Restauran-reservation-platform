import './RestaurantPage.css'

import { useState, useEffect } from 'react'

import { Link } from 'react-router-dom'

import Header from '../../containers/Header/Header'
import Footer from '../../containers/Footer/Footer'

import capitalizeFirstLetterOfEachWord from '../../helpers/capitalizeFunction.js'
import IntToDay from '../../helpers/intToDay.js'

import GetRestaurantById from '../../api/GetRestaurantById'

export default function RestaurantPage() {
  let restaurant_id = String(window.location.href)
  restaurant_id = restaurant_id.split("/")
  restaurant_id = restaurant_id[restaurant_id.length - 1]

  const [restaurant, setRestaurant] = useState([])

  const [daysOpen, setDaysOpen] = useState([])
  const [workingHours, setWorkingHours] = useState([])
  const [Specialties, setSpecialties] = useState([])

  useEffect(() => {
    const fetchRestaurant = async () => {
      const restaurantResponse = await GetRestaurantById(restaurant_id)
      setRestaurant(restaurantResponse)

      setDaysOpen(restaurantResponse.days_open)
      setWorkingHours(restaurantResponse.working_hours)
      setSpecialties(restaurantResponse.specialties)
    }

    fetchRestaurant()
  }, [])

  return (
    <>
      <Header />

      <main className='Restaurantpage-main'>
        <section className='Restaurantpage-section'>
          <h3>{capitalizeFirstLetterOfEachWord(restaurant.name)}</h3>
          <p><i>{restaurant.id}</i></p>
          <p>{capitalizeFirstLetterOfEachWord(restaurant.city) + "."}</p>
          <p>{capitalizeFirstLetterOfEachWord(restaurant.address) + "."}</p>
          <p>{restaurant.description}</p>
          <p>{daysOpen.map((day, index) => { if(index == daysOpen.length - 1) { return (IntToDay(day) + ".") } else { return (IntToDay(day) + ", ") } })}</p>
          <p>{workingHours.map((wh, index) => { if(index % 2 == 0) { return (wh + "-") } else { return (wh + ". ") } })}</p>
          <p>{Specialties.map((speciality, index) => { if(index == Specialties.length - 1) { return (capitalizeFirstLetterOfEachWord(speciality) + ".") } else { return (capitalizeFirstLetterOfEachWord(speciality) + ", ") } })}</p>
          <p className='Reservation-button'><Link to={`book`}>Reservar</Link></p>
        </section>
      </main>

      <Footer />
    </>
  )
}

