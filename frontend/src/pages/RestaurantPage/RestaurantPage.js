import './RestaurantPage.css'

import { Link } from 'react-router-dom'

import Header from '../../containers/Header/Header'
import Footer from '../../containers/Footer/Footer'

export default function RestaurantPage() {

  const setRestaurantId = () => {
    localStorage.setItem('restaurant_id', 'id')
  }

  return (
    <>
      <Header />

      <main className='Restaurantpage-main'>
        <section className='Restaurantpage-section'>
          <h3>Nombre del restaurante</h3>
          <p><i>id</i></p>
          <p>Ciudad</p>
          <p>Dirección</p>
          <p>Descripción: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer ultrices purus quis cursus congue. Integer in urna pharetra, lacinia quam id, cursus est. Donec eleifend, velit scelerisque volutpat imperdiet, neque dui rutrum turpis, a euismod tortor augue in ligula. Integer ac odio posuere, iaculis est et, cursus nisl. Ut malesuada.</p>
          <p>Día 1, Día 2, Día 3, Día 4</p>
          <p>[20:00, 24:00], [20:00, 24:00], [17:00, 24:00], [17:00, 24:00]</p>
          <p>Especialidad 1, Especialidad 2, Especialidad 3</p>
          <p className='Reservation-button'><Link onClick={setRestaurantId} to="/new-reservation">Reservar</Link></p>
        </section>
      </main>

      <Footer />
    </>
  )
}

