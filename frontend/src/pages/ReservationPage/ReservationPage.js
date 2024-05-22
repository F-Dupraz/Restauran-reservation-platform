import './ReservationPage.css'

import { useRef } from 'react'

import Header from '../../containers/Header/Header'
import Footer from '../../containers/Footer/Footer'

export default function ReservationPage() {

  const from_h = useRef(null)
  const to_h = useRef(null)
  const day = useRef(null)
  const num_guests = useRef(null)

  const submitData = () => {
    if(!from_h.current.value || !to_h.current.value || !day.current.value || !num_guests.current.value) {
      alert("Tienes que completar todos los campos si quieres editar la reserva.")
    } else {
      const data = {
        from_h: from_h.current.value,
        to_h: to_h.current.value,
        day: day.current.value,
        num_guests: num_guests.current.value
      }
      console.log(data)
    }
  }

  const deleteReservation = () => {
    console.log("Reservation Deleted")
  }

  const displayForm = () => {
    const form_section = document.querySelector(".reservationpage-form-div")
    form_section.style.top = "25%"
    form_section.style.left = "50%"
  }

  return (
    <>
    <Header />

    <main className='reservationpage-main'>
      <section className='reservationpage-section'>
        <div>
          <h3>Restaurante</h3>
          <p>01/01/2001</p>
          <p>Desde: 20:00hs</p>
          <p>Hasta: 23:00hs</p>
          <p>Para 5 personas</p>
          <input type='button' name='edit' id='edit-reservation-button' onClick={displayForm} value="Editar" />
          <input type="button" id="delete-reservation-button" onClick={deleteReservation} value="Eliminar Reserva" />
        </div>
        <div className='reservationpage-form-div'>
          <form>
            <label>
              Desde:
              <input type="time" name="from_h" id="from_h" required={true} ref={from_h} />
            </label>
            <label>
              Hasta:
              <input type="time" name="to_h" id="to_h" required={true} ref={to_h} />
            </label>
            <label>
              Día:
              <input type="date" name="day" id="day" required={true} ref={day} />
            </label>
            <label>
              Cantidad:
              <input type="number" name="num_guests" id="num_guests" required={true} ref={num_guests} />
            </label>
            <input type="button" id="reservation-button" onClick={submitData} value="Editar" />
          </form>
        </div>
      </section>
    </main>

    <Footer />
    </>
  )
}
