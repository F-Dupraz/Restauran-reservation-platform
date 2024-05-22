import './CreateReservation.css'

import { useRef } from 'react'

import Header from '../../containers/Header/Header'
import Footer from '../../containers/Footer/Footer'

export default function CreateReservation() {
  const date = useRef(null)
  const start_time = useRef(null)
  const end_time = useRef(null)
  const num_guests = useRef(null)

  const submitData = () => {
    if(!date.current.value || !start_time.current.value || !end_time.current.value || !num_guests.current.value) {
      alert("Tienes que completar todos los campos si quieres crear un usuario.")
    } else {
      const data = {
        user_id: 'id',
        restaurant_id: localStorage.getItem('restaurant_id'),
        date: date.current.value,
        start_time: start_time.current.value,
        end_time: end_time.current.value,
        num_guests: num_guests.current.value,
      }
      console.log(data)
    }
  }

  return (
    <>
    <Header />

    <main className='createreservation-main'>
      <section className='createreservation-section'>
        <h3>Nueva Reserva</h3>
        <form className="createreservation-form" method="post">
          <label>
            Día:
            <input type='date' name="date" id="date" required={true} ref={date} />
          </label>
          <label>
            Horario de inicio:
            <input type="time" name="start_time" id="start_time" required={true} ref={start_time} />
          </label>
          <label>
            Horario de salida:
            <input type="time" name="end_time" id="end_time" required={true} ref={end_time} />
          </label>
          <label>
            Número de invitados:
            <input type="number" name="num_guests" id="num_guests" required={true} ref={num_guests} />
          </label>
          <input type="button" id="res-button" onClick={submitData} value="Reservar" />
        </form>
      </section>
    </main>

    <Footer />
    </>
  )
}
