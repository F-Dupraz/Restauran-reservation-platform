import './ReservationPage.css'

import { useRef, useState, useEffect } from 'react'

import { useNavigate } from 'react-router-dom'

import Header from '../../containers/Header/Header'
import Footer from '../../containers/Footer/Footer'

import capitalizeFirstLetterOfEachWord from '../../helpers/capitalizeFunction.js'

import GetReservationById from '../../api/GetReservationById'
import EditReservation from '../../api/EditReservation'
import DeleteReservation from '../../api/DeleteReservation'

export default function ReservationPage() {
  const navigation = useNavigate()

  const [name, setReservationName] = useState(null)
  const [from, setReservationFrom] = useState(null)
  const [to, setReservationTo] = useState(null)
  const [day, setReservationDay] = useState(null)
  const [numGuests, setReservationNumGuests] = useState(null)

  const my_token = localStorage.getItem("my_token")

  let reservation_id = String(window.location.href)
  reservation_id = reservation_id.split("/")
  reservation_id = reservation_id[reservation_id.length - 1]

  const from_h = useRef(null)
  const to_h = useRef(null)
  const day_h = useRef(null)
  const num_guests = useRef(null)

  const submitData = async () => {
    if(!from_h.current.value || !to_h.current.value || !day_h.current.value || !num_guests.current.value) {
      alert("Tienes que completar todos los campos si quieres editar la reserva.")
    } else {
      const data = {
        id: reservation_id,
        from: from_h.current.value,
        to: to_h.current.value,
        day: day_h.current.value,
        num_guests: parseInt(num_guests.current.value)
      }

      const response = await EditReservation(data, my_token)
      if(response.success) {
        alert(response.message)
        navigation("/userpage")
      }
    }
  }

  const deleteReservation = async () => {
    const response = await DeleteReservation(reservation_id, my_token)
    if(response.success) {
      alert(response.message)
      navigation("/userpage")
    }
  }

  const displayForm = () => {
    const form_section = document.querySelector(".reservationpage-form-div")
    form_section.style.top = "25%"
    form_section.style.left = "50%"
  }

  useEffect(() => {
    const fetchReservation = async () => {
      try {
        const reservationResponse = await GetReservationById(reservation_id, my_token)
        setReservationName(reservationResponse.restaurant_name)
        setReservationDay(reservationResponse.day)
        setReservationFrom(reservationResponse.from)
        setReservationTo(reservationResponse.to)
        setReservationNumGuests(reservationResponse.num_guests)
      } catch(err) {
        console.log(err)
      }
    }
  
    fetchReservation()
  }, [])

  return (
    <>
    <Header />

    <main className='reservationpage-main'>
      <section className='reservationpage-section'>
        <div>
          <h3>{capitalizeFirstLetterOfEachWord(name)}</h3>
          <p>{day == null ? "" : day.slice(0, 10)}</p>
          <p>Desde: {from == null ? "" : from.slice(11, -4)}</p>
          <p>Hasta: {to == null ? "" : to.slice(11, -4)}</p>
          <p>Para {numGuests} personas</p>
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
              <input type="date" name="day" id="day" required={true} ref={day_h} />
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
