import './HomePage.css'

import { Link, useNavigate } from 'react-router-dom'

import isJwtExpired from '../../helpers/getJWT'

import Header from '../../containers/Header/Header'
import Footer from '../../containers/Footer/Footer'
import { useEffect } from 'react'

export default function HomePage() {
  const navigate = useNavigate()

  useEffect(() => {
    const my_token = localStorage.getItem('my_token')
    
    if(my_token && !isJwtExpired(my_token)) {
      navigate("/restaurants")
    }
  }, [navigate])

  return (
    <>
    <Header />

    <main className='main-section'>
      <section className='text-section'>
        <h2>Descubre, Reserva y Disfruta: MesaBook, Tu Guía Gastronómica Definitiva</h2>
        <p>
          ¡Bienvenido a MesaBook, tu aplicación definitiva para descubrir y reservar en los mejores restaurantes! Con MesaBook, explorar la escena culinaria local nunca ha sido tan fácil y emocionante.<br /><br />
          Imagina tener acceso a una amplia selección de restaurantes auténticos, desde acogedores cafés hasta lujosos restaurantes de cinco estrellas, todo al alcance de tu mano.<br /><br />
          Nos enorgullecemos de ofrecer a nuestros usuarios solo los mejores establecimientos gastronómicos, para que cada comida sea una experiencia inolvidable.<br /><br />
          ¿Quieres asegurarte de tener una mesa reservada en tu restaurante favorito? ¡MesaBook lo hace posible en solo unos pocos clics! Con nuestra función de reserva fácil de usar, puedes garantizar tu lugar en el restaurante que elijas sin complicaciones ni esperas. Ya sea para una cena romántica, una reunión de negocios o una comida casual con amigos, MesaBook te ofrece la tranquilidad de saber que tu mesa está lista cuando llegues.<br /><br />
          Únete a la comunidad de amantes de la comida que confían en MesaBook para sus aventuras gastronómicas. Descarga la aplicación hoy mismo y comienza a explorar, reservar y disfrutar de los mejores restaurantes cerca de ti.<br /><br />
          ¡Buen provecho!<br />
        </p>
      </section>
      <section className='create-account-section'>
        <h3>¿Aún no tienes una cuenta?</h3>
        <div>
          <Link to="create-account">Crear Cuenta</Link>
        </div>
      </section>
    </main>
    
    <Footer />
    </>
  )
}
