import utils from '../utils'

// If the url of the dog object is a video return a video component otherwise return an image
function Media(props){
    return utils.isVideo(props.dog.url) ? 
    (
        <video className="home-img" width="320" height="240" controls>
            <source src={props.dog.url} type="video/mp4" alt="dogs"/>
            <source src={props.dog.url} type="video/ogg" alt="dogs"/>
            Your browser does not support the video tag.
        </video>
    )
    :
    (
        <img className="home-img" src={props.dog.url} loading="lazy" alt="dogs"/>
    )
}

export default Media