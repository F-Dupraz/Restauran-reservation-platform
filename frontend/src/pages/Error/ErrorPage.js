import './ErrorPage.css'

import { Link, useRouteError  } from 'react-router-dom'

import Header from '../../containers/Header/Header'

export default function ErrorPage() {
  const error = useRouteError()

  return (
    <>
    <Header />

    <main className='errorpage-main'>
      <section className='errorpage-section'>
        <h3>¡Ocurrió un error!</h3>
        <p>Lamentamos que haya ocurido este error.</p>
        <p>
          <i>{ error.message || error.statusText }</i>
        </p>
        <p>
          <Link to="/">Vuelve al inicio</Link>
        </p>
      </section>
    </main>
    </>
  )
}
