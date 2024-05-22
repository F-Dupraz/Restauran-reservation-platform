import './EditRestaurantPage.css'

import { useRef } from 'react'

import Header from '../../containers/Header/Header'
import Footer from '../../containers/Footer/Footer'

export default function EditRestaurantPage() {
  const name = useRef(null)
  const description = useRef(null)

  const submitData = () => {
    console.log("Submit restaurant")
  }

  const deleteRestaurant = () => {
    console.log("Delet restaurant")
  }

  return (
    <>
    <Header />

    <main className='editrestaurant-main'>
      <section className='editrestaurant-section'>
        <h3>Editar Restaurante</h3>
        <form className='editrestaurant-form'>
          <label>
            Nombre:
            <input type="text" name="name" id="edit-name" required={true} ref={name} />
          </label>
           <label>
            Description:
            <input type="text" name="description" id="edit-description" required={true} ref={description} />
          </label>
          <label className='createrestaurant-days'>
            Días abierto:<br />
            <div>
              <div>
                <input type="checkbox" name="monday" id="edit-monday" /><label>Lunes</label>
              </div>
              <div>
                <input type="checkbox" name="tuesday" id="edit-tuesday" /><label>Martes</label>
              </div>
              <div>
                <input type="checkbox" name="wednesday" id="edit-wednesday" /><label>Miércoles</label>
              </div>
              <div>
                <input type="checkbox" name="thursday" id="edit-thursday" /><label>Jueves</label>
              </div>
              <div>
                <input type="checkbox" name="friday" id="edit-friday" /><label>Viernes</label>
              </div>
              <div>
                <input type="checkbox" name="saturday" id="edit-saturday" /><label>Sábado</label>
              </div>
              <div>
                <input type="checkbox" name="sunday" id="edit-sunday" /><label>Domingo</label>
              </div>
            </div>
          </label>
          <label className='createrestaurant-h'>
            Horarios de apertura:
            <div>
              <input type="time" name="h_monday" id="edit-h_monday" />
              <input type="time" name="h_tuesday" id="edit-h_tuesday" />
              <input type="time" name="h_wednesday" id="edit-h_wednesday" />
              <input type="time" name="h_thursday" id="edit-h_thursday" />
              <input type="time" name="h_friday" id="edit-h_friday" />
              <input type="time" name="h_saturday" id="edit-h_saturday" />
              <input type="time" name="h_sunday" id="edit-h_sunday" />
            </div>
          </label>
          <label className='createrestaurant-h_h'>
            Horarios de cierre:
            <div>
              <input type="time" name="h_monday_h" id="edit-h_monday_h" />
              <input type="time" name="h_tuesday_h" id="edit-h_tuesday_h" />
              <input type="time" name="h_wednesday_h" id="edit-h_wednesday_h" />
              <input type="time" name="h_thursday_h" id="edit-h_thursday_h" />
              <input type="time" name="h_friday_h" id="edit-h_friday_h" />
              <input type="time" name="h_saturday_h" id="edit-h_saturday_h" />
              <input type="time" name="h_sunday_h" id="edit-h_sunday_h" />
            </div>
          </label>
          <label className='createrestaurant-capacity'>
            Capacidades por día:
            <div>
              <input type='number' name='c_monday' id='edit-c_monday'/>
              <input type='number' name='c_tuesday' id='edit-c_tuesday'/>
              <input type='number' name='c_wednesday' id='edit-c_wednesday'/>
              <input type='number' name='c_thursday' id='edit-c_thursday'/>
              <input type='number' name='c_friday' id='edit-c_friday'/>
              <input type='number' name='c_satuday' id='edit-c_satuday'/>
              <input type='number' name='c_sunday' id='edit-c_sunday'/>
            </div>
          </label>
          <label>
            Especialidades:
            <input type='text' name='specialities' id='edit-specialities' required={true} />
          </label>
          <input type='button' className='submit-button' onClick={submitData} name='createrestaurant-button' id='createrestaurant-button' value='Editar Restaurante'/><br />
          <input type='button' className='delete-button' onClick={deleteRestaurant} name='deletetaurant-button' id='deletetaurant-button' value='Eliminar Restaurante'/> 
        </form>
      </section>
    </main>

    <Footer />
    </>
  )
}
