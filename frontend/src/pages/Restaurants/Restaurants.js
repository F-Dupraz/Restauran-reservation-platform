import './Restaurants.css'

import { useState, useEffect } from 'react'

import { Link } from 'react-router-dom'

import Header from '../../containers/Header/Header'
import Footer from '../../containers/Footer/Footer'

import capitalizeFirstLetterOfEachWord from '../../helpers/capitalizeFunction.js'

import GetRestaurants from '../../api/GetRestaurants.js'

export default function Restaurants() {
  let offsetParam = 0

  const [restaurants, setRestaurants] = useState([])

  const handleScroll = async () => {
    if((document.documentElement.scrollTop + document.documentElement.clientHeight) >= (document.documentElement.scrollHeight - 80)) {
      offsetParam += 20
      const restaurantsResponse = await GetRestaurants(offsetParam)
      setRestaurants(prevRestaurants => [...prevRestaurants, ...restaurantsResponse])
    }
  }

  useEffect(() => {
    const fetchRestaurants = async () => {
      const restaurantsResponse = await GetRestaurants(offsetParam)
      setRestaurants(restaurantsResponse)
    }

    fetchRestaurants()

    document.addEventListener("scroll", handleScroll)
  }, [])

  return (
    <>
      <Header />

      <main>
        <section className='restaurant-section'>
          {
            restaurants.map((restaurant) => {
              return (
                <div className='restaurant-element' key={restaurant.id}>
                  <Link to={`${restaurant.id}`}>
                    <h4>{capitalizeFirstLetterOfEachWord(restaurant.name)}</h4>
                    <p>{capitalizeFirstLetterOfEachWord(restaurant.city)}</p>
                    <p>{capitalizeFirstLetterOfEachWord(restaurant.address)}</p>
                    <div className='restaurant-element-specialties'>
                      {
                        restaurant.specialties.map((speciality, index) => {
                          if(index == restaurant.specialties.length - 1) {
                            return <p key={index}>{capitalizeFirstLetterOfEachWord(speciality) + "."}</p>
                          } else {
                            return <p key={index}>{capitalizeFirstLetterOfEachWord(speciality) + ", "}</p>
                          }
                        })
                      }
                    </div>
                    <p><i>{restaurant.id}</i></p>
                  </Link>
                </div>
              )
            })
          }
        </section>
      </main>

      <Footer /> 
    </>
  )
}
