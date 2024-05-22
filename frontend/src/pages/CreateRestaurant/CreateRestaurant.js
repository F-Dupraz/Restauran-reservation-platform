import './CreateRestaurant.css'

import { useRef } from 'react'

import Header from '../../containers/Header/Header'
import Footer from '../../containers/Footer/Footer'

export default function CreateRestaurant() {
  const name = useRef(null)
  const city = useRef(null)
  const address = useRef(null)
  const description = useRef(null)

  const submitData = () => {
    console.log("Submit restaurant")
  }

  return (
    <>
    <Header />

    <main className='createrestaurant-main'>
      <section className='createrestaurant-section'>
      <h3>Nuevo Restaurante</h3>
      <label>
        Nombre:
        <input type="text" name="name" id="name" required={true} ref={name} />
      </label>
      <label>
        Ciudad:
        <input type="text" name="city" id="city" required={true} ref={city} />
      </label>
      <label>
        Dirección:
        <input type="text" name="address" id="address" required={true} ref={address} />
      </label>
      <label>
        Description:
        <input type="text" name="description" id="description" required={true} ref={description} />
      </label>
      <label className='createrestaurant-days'>
        Días abierto:<br />
        <div>
          <div>
            <input type="checkbox" name="monday" id="monday" /><label>Lunes</label>
          </div>
          <div>
            <input type="checkbox" name="tuesday" id="tuesday" /><label>Martes</label>
          </div>
          <div>
            <input type="checkbox" name="wednesday" id="wednesday" /><label>Miércoles</label>
          </div>
          <div>
            <input type="checkbox" name="thursday" id="thursday" /><label>Jueves</label>
          </div>
          <div>
            <input type="checkbox" name="friday" id="friday" /><label>Viernes</label>
          </div>
          <div>
            <input type="checkbox" name="saturday" id="saturday" /><label>Sábado</label>
          </div>
          <div>
            <input type="checkbox" name="sunday" id="sunday" /><label>Domingo</label>
          </div>
        </div>
      </label>
      <label className='createrestaurant-h'>
        Horarios de apertura:
        <div>
          <input type="time" name="h_monday" id="h_monday" />
          <input type="time" name="h_tuesday" id="h_tuesday" />
          <input type="time" name="h_wednesday" id="h_wednesday" />
          <input type="time" name="h_thursday" id="h_thursday" />
          <input type="time" name="h_friday" id="h_friday" />
          <input type="time" name="h_saturday" id="h_saturday" />
          <input type="time" name="h_sunday" id="h_sunday" />
        </div>
      </label>
      <label className='createrestaurant-h_h'>
        Horarios de cierre:
        <div>
          <input type="time" name="h_monday_h" id="h_monday_h" />
          <input type="time" name="h_tuesday_h" id="h_tuesday_h" />
          <input type="time" name="h_wednesday_h" id="h_wednesday_h" />
          <input type="time" name="h_thursday_h" id="h_thursday_h" />
          <input type="time" name="h_friday_h" id="h_friday_h" />
          <input type="time" name="h_saturday_h" id="h_saturday_h" />
          <input type="time" name="h_sunday_h" id="h_sunday_h" />
        </div>
      </label>
      <label className='createrestaurant-capacity'>
        Capacidades por día:
        <div>
          <input type='number' name='c_monday' id='c_monday'/>
          <input type='number' name='c_tuesday' id='c_tuesday'/>
          <input type='number' name='c_wednesday' id='c_wednesday'/>
          <input type='number' name='c_thursday' id='c_thursday'/>
          <input type='number' name='c_friday' id='c_friday'/>
          <input type='number' name='c_satuday' id='c_satuday'/>
          <input type='number' name='c_sunday' id='c_sunday'/>
        </div>
      </label>
      <label>
        Especialidades:
        <input type='text' name='specialities' id='specialities' required={true} />
      </label>
      <input type='button' className='submit-button' onClick={submitData} name='createrestaurant-button' id='createrestaurant-button' value='Crear Restaurante' /> 
      </section>
    </main>

    <Footer />
    </>
  )
}