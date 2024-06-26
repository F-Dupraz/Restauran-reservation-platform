import './CreateRestaurant.css'

import { useRef } from 'react'
import { useNavigate } from 'react-router-dom'

import Header from '../../containers/Header/Header'
import Footer from '../../containers/Footer/Footer'

import PostRestaurant from '../../api/PostRestaurant'

export default function CreateRestaurant() {
  const my_token = localStorage.getItem("my_token")

  const navigation = useNavigate()

  const name = useRef(null)
  const city = useRef(null)
  const address = useRef(null)
  const description = useRef(null)

  const monday = useRef(null)
  const tuesday = useRef(null)
  const wednesday = useRef(null)
  const thursday = useRef(null)
  const friday = useRef(null)
  const saturday = useRef(null)
  const sunday = useRef(null)
  const days = [monday, tuesday, wednesday, thursday, friday, saturday, sunday]

  const h_monday = useRef(null)
  const h_tuesday = useRef(null)
  const h_wednesday = useRef(null)
  const h_thursday = useRef(null)
  const h_friday = useRef(null)
  const h_saturday = useRef(null)
  const h_sunday = useRef(null)
  const h_days = [h_monday, h_tuesday, h_wednesday, h_thursday, h_friday, h_saturday, h_sunday]

  const h_monday_h = useRef(null)
  const h_tuesday_h = useRef(null)
  const h_wednesday_h = useRef(null)
  const h_thursday_h = useRef(null)
  const h_friday_h = useRef(null)
  const h_saturday_h = useRef(null)
  const h_sunday_h = useRef(null)
  const h_days_h = [h_monday_h, h_tuesday_h, h_wednesday_h, h_thursday_h, h_friday_h, h_saturday_h, h_sunday_h]

  const c_monday = useRef(null)
  const c_tuesday = useRef(null)
  const c_wednesday = useRef(null)
  const c_thursday = useRef(null)
  const c_friday = useRef(null)
  const c_saturday = useRef(null)
  const c_sunday = useRef(null)
  const c_days = [c_monday, c_tuesday, c_wednesday, c_thursday, c_friday, c_saturday, c_sunday]

  const specialities = useRef(null)

  const submitData = async () => {
    let days_open = []
    let from = []
    let to = []
    let capacity = []

    let specialities_parsed = specialities.current.value
    specialities_parsed = specialities_parsed.split(",")
    for(let i = 0; i < specialities_parsed.length; i++) {
      specialities_parsed[i] = specialities_parsed[i].trim()
    }

    for(let i = 0; i < 7; i++) {
      if(days[i].current.checked && h_days[i].current.value != "" && h_days_h[i].current.value != "" && c_days[i].current.value != "") {
        days_open.push(parseInt(days[i].current.value))
        from.push(h_days[i].current.value)
        to.push(h_days_h[i].current.value)
        capacity.push(parseInt(c_days[i].current.value))
      } else {
        continue
      }
    }

    if(!name.current.value || !city.current.value || !address.current.value || !description.current.value) {
      alert("Tienes que completar todos los campos si quieres crear un restaurante.")
    } else if(days_open.length <= 0 || from.length <= 0 || to.length <= 0 || capacity.length <= 0) {
      alert("Tienes que completar todos los campos si quieres crear un restaurante.")
    } else {
      const data = {
        name: name.current.value,
        city: city.current.value,
        address: address.current.value,
        description: description.current.value,
        days_open: days_open,
        from: from,
        to: to,
        capacity: capacity,
        specialties: specialities_parsed
      }
  
      const response = await PostRestaurant(data, my_token)
      if(response.success) {
        alert(response.message)
        navigation("/restaurants")
      }
    } 
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
            <input type="checkbox" name="monday" id="monday" value="1" ref={monday}/><label>Lunes</label>
          </div>
          <div>
            <input type="checkbox" name="tuesday" id="tuesday" value="2" ref={tuesday}/><label>Martes</label>
          </div>
          <div>
            <input type="checkbox" name="wednesday" id="wednesday" value="3" ref={wednesday}/><label>Miércoles</label>
          </div>
          <div>
            <input type="checkbox" name="thursday" id="thursday" value="4" ref={thursday}/><label>Jueves</label>
          </div>
          <div>
            <input type="checkbox" name="friday" id="friday" value="5" ref={friday}/><label>Viernes</label>
          </div>
          <div>
            <input type="checkbox" name="saturday" id="saturday" value="6" ref={saturday}/><label>Sábado</label>
          </div>
          <div>
            <input type="checkbox" name="sunday" id="sunday" value="0" ref={sunday}/><label>Domingo</label>
          </div>
        </div>
      </label>
      <label className='createrestaurant-h'>
        Horarios de apertura:
        <div>
          <input type="time" name="h_monday" id="h_monday" ref={h_monday}/>
          <input type="time" name="h_tuesday" id="h_tuesday" ref={h_tuesday}/>
          <input type="time" name="h_wednesday" id="h_wednesday" ref={h_wednesday}/>
          <input type="time" name="h_thursday" id="h_thursday" ref={h_thursday}/>
          <input type="time" name="h_friday" id="h_friday" ref={h_friday}/>
          <input type="time" name="h_saturday" id="h_saturday" ref={h_saturday}/>
          <input type="time" name="h_sunday" id="h_sunday" ref={h_sunday}/>
        </div>
      </label>
      <label className='createrestaurant-h_h'>
        Horarios de cierre:
        <div>
          <input type="time" name="h_monday_h" id="h_monday_h" ref={h_monday_h}/>
          <input type="time" name="h_tuesday_h" id="h_tuesday_h" ref={h_tuesday_h}/>
          <input type="time" name="h_wednesday_h" id="h_wednesday_h"ref={h_wednesday_h} />
          <input type="time" name="h_thursday_h" id="h_thursday_h" ref={h_thursday_h}/>
          <input type="time" name="h_friday_h" id="h_friday_h" ref={h_friday_h}/>
          <input type="time" name="h_saturday_h" id="h_saturday_h" ref={h_saturday_h}/>
          <input type="time" name="h_sunday_h" id="h_sunday_h" ref={h_sunday_h}/>
        </div>
      </label>
      <label className='createrestaurant-capacity'>
        Capacidades por día:
        <div>
          <input type='number' name='c_monday' id='c_monday' ref={c_monday} />
          <input type='number' name='c_tuesday' id='c_tuesday' ref={c_tuesday} />
          <input type='number' name='c_wednesday' id='c_wednesday' ref={c_wednesday} />
          <input type='number' name='c_thursday' id='c_thursday' ref={c_thursday} />
          <input type='number' name='c_friday' id='c_friday' ref={c_friday} />
          <input type='number' name='c_satuday' id='c_satuday' ref={c_saturday} />
          <input type='number' name='c_sunday' id='c_sunday' ref={c_sunday} />
        </div>
      </label>
      <label>
        Especialidades:
        <input type='text' name='specialities' id='specialities' required={true} ref={specialities} />
      </label>
      <input type='button' className='submit-button' onClick={submitData} name='createrestaurant-button' id='createrestaurant-button' value='Crear Restaurante' /> 
      </section>
    </main>

    <Footer />
    </>
  )
}