import './Footer.css'

import { Link } from 'react-router-dom'

export default function Footer() {
  return (
    <footer id="footer">
      <div className='footer-links'>
        <ul>
          <li>Mail: duprazfabricio@gmail.com</li>
          <li><Link to="/terms-and-conditions">Terminos y condiciones</Link></li>
          <li><a href="https://x.com/fabridupraz" target="_blank" rel="noopener noreferrer">Twitter</a></li>
        </ul>
      </div>
      <div className='footer-thanks'>
        <p>¡Gracias por todo!</p>
      </div>
    </footer>
  )
}