import './App.css'

// Importações de componentes
import Header from './assets/components/Header'
import Carousel from './assets/components/Carousel'
import Product from './assets/components/Product'

function App() {

  return (
    <body className="bg-custom-background">

      <div className='items-center space-y-2'>
        <Header/>

        <Carousel/>

        <div>
          <p>Ganhe um desconto usando o cupom "SMART"</p>
        </div>

        {/* Adicionar dentro do banco os produtos e criar requisições */}
        <div>
          <Product Name="Halter"></Product>
        </div>

      </div>

    </body>
  )
}

export default App
