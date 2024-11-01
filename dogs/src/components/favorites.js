import utils from '../utils'
import Media from './media'

function Favorites(){
    const dogs = utils.getDogs().map((val, index)=>{
        return <li key={index}>
            <Media dog={val}/>
        </li>
    })

    return (<div className="favorites">
        <h1>Hi! ✨ Lets see your dogs 😃</h1>
        <ul className="dogs-container shadow">
            {dogs}
        </ul>
    </div>)
}

export default Favorites