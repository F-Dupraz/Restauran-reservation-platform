export default function IntToDay(IntDay) {
  let StrDay = ""
    switch(IntDay) {
      case 0:
        StrDay = "Domingo"
        break
      case 1:
        StrDay = "Lunes"
        break
      case 2:
        StrDay = "Martes"
        break
      case 3:
        StrDay = "Miércoles"
        break
      case 4:
        StrDay = "Jueves"
        break
      case 5:
        StrDay = "Viernes"
        break
      case 6:
        StrDay = "Sábado"
        break
    }
  return StrDay
}