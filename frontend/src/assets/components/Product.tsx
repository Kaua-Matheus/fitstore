
// Temos que inserir as props para realmente poder passar parâmetros
interface ProductProps {
    Name: string
}

export default function Product({Name}: ProductProps) {

    return (
        <div>
            {`${Name}`}
        </div>
    )
}