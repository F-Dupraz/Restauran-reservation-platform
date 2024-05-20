import './Header.css'

import { useRef, useState } from 'react'

import { Link } from 'react-router-dom'

export default function Header() {
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

  const [styles, setStyles] = useState({
    top: '-100%',
    transition: '0s'
  })

  const displayLogin = () => {
    setStyles({
      top: '50%',
      transition: '1s ease-out'
    })
  }

  return (
    <>
      <div className='login-section' style={styles}>
        <h3>Inicia Sesión</h3>
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
      </div>

      <header>
        <div className='header-title-div'>
          <h1>MesaBook</h1>
        </div>
        <div className='header-login-div'>
          <Link onClick={displayLogin}>Log In</Link>
        </div>
      </header>
    </>
  )
}
