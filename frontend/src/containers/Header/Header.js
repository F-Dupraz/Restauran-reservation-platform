import './Header.css'

import { Link } from 'react-router-dom'

export default function Header() {
  return (
    <header>
      <div className='header-title-div'>
        <h1>MesaBook</h1>
      </div>
      <div className='header-login-div'>
        <Link to="/login">Log In</Link>
      </div>
    </header>
  )
}
