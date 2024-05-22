import './UserPage.css'

import { Link } from 'react-router-dom'

import Header from '../../containers/Header/Header'
import Footer from '../../containers/Footer/Footer'

export default function UserPage() {
  return (
    <>
      <Header />

      <main>
        <section className='userpage-section'>
          <div className='userpage-info'>
            <h2>Username</h2>
            <p>ejemplo@mail.cmo</p>
            <p>id</p>
          </div>
          <div className='userpage-restaurants'>
            <h3>Mis Restaurantes</h3>
            <div className='userpage-restaurant-container'>
              <div className='userpage-restaurant'>
                <Link to="/restaurants/:id">
                  <h4>Nombre</h4>
                  <p><i>id</i></p>
                </Link>
                <Link to="/edit-restaurants/:id" className='editrestaurant-link'>
                  Editar
                </Link>
              </div>
              <div className='userpage-restaurant'>
                <Link to="/restaurants/:id">
                  <h4>Nombre</h4>
                  <p><i>id</i></p>
                </Link>
                <Link to="/edit-restaurants/:id" className='editrestaurant-link'>
                  Editar
                </Link>
              </div>
              <div className='userpage-restaurant'>
                <Link to="/restaurants/:id">
                  <h4>Nombre</h4>
                  <p><i>id</i></p>
                </Link>
                <Link to="/edit-restaurants/:id" className='editrestaurant-link'>
                  Editar
                </Link>
              </div>
            </div>
            <p className='userpage-restaurants-p'>
              <Link to="/new-restaurant">Añadir Restaurante</Link>
            </p>
          </div>
          <div className='userpage-reservations'>
            <h3>Mis Reservas</h3>
            <div className='userpage-reservation'>
              <Link to="/reservations/:id">
                <p>Restaurante</p>
                <p>01/01/2001</p>
                <p><i>id</i></p>
              </Link>
            </div>
            <div className='userpage-reservation'>
              <Link to="/reservations/:id">
                <p>Restaurante</p>
                <p>01/01/2001</p>
                <p><i>id</i></p>
              </Link>
            </div>
            <div className='userpage-reservation'>
              <Link to="/reservations/:id">
                <p>Restaurante</p>
                <p>01/01/2001</p>
                <p><i>id</i></p>
              </Link>
            </div>
            <div className='userpage-reservation'>
              <Link to="/reservations/:id">
                <p>Restaurante</p>
                <p>01/01/2001</p>
                <p><i>id</i></p>
              </Link>
            </div>
          </div>
        </section>
      </main>

      <Footer />
    </>
  )
}