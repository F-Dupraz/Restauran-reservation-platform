import './CreateUser.css'

import { useRef } from 'react'

import { Link } from 'react-router-dom'

import Header from '../../containers/Header/Header'
import Footer from '../../containers/Footer/Footer'

export default function CreateUser() {
  const username = useRef(null)
  const email = useRef(null)
  const password = useRef(null)

  const submitData = () => {
    if(!username.current.value || !email.current.value || !password.current.value) {
      alert("Tienes que completar todos los campos si quieres crear un usuario.")
    } else {
      const data = {
        username: username.current.value,
        email: email.current.value,
        password: password.current.value
      }
      console.log(data)
    }
  }

  return (
    <>
    <Header />

    <main className='createuser-main'>
      <section className='createuser-section'>
        <form className="createuser-form" method="post">
          <label>
            Username
            <input type="text" name="username" id="username" required={true} ref={username} />
          </label>
          <label>
            Email
            <input type="email" name="email" id="email" required={true} ref={email} />
          </label>
          <label>
            Password
            <input type="password" name="password" id="password" required={true} ref={password} />
          </label>
          <input type="button" id="button" onClick={submitData} value="Crear usuario" />
        </form>

        <p>
          <Link to="/login">¿Ya tienes un usuario?</Link>
        </p>
      </section>
    </main>

    <Footer />
    </>
  )
}
