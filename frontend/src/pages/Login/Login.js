import './Login.css'

import { useRef } from 'react'

import { Link } from 'react-router-dom'

import Header from '../../containers/Header/Header'
import Footer from '../../containers/Footer/Footer'

export default function Login() {
  const email = useRef(null)
  const password = useRef(null)

  const submitData = () => {
    if(!email.current.value || !password.current.value) {
      alert("Tienes que completar todos los campos si quieres iniciar sesión.")
    } else {
      const data = {
        email: email.current.value,
        password: password.current.value
      }
      console.log(data)
    }
  }

  return (
    <>
    <Header />

    <main className='login-main'>
      <section className='login-leftsection'>
        <p>
          Descubre los mejores restaurantes y asegúrate de tener tu mesa reservada y lista cuando llegues.<br /><br />
          Con MesaBook, explorar y reservar es rápido y fácil.<br /><br />
          Inicia sesión ahora y comienza tu experiencia gastronómica sin preocupaciones.<br /><br />
        </p>
        <form>
          <label>
            Email:
            <input type="email" name="email" id="email" required={true} ref={email} />
          </label>
          <label>
            Password:
            <input type="password" name="password" id="password" required={true} ref={password} />
          </label>
          <input type="button" id="login-button" onClick={submitData} value="Iniciar Sesión" />
        </form>
        <p>
          <i>¿Nuevo en MesaBook? <Link to="/create-account" >Registrate aquí.</Link></i>
        </p>
      </section>
    </main>
    
    <Footer />
    </>
  )
}