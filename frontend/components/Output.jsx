import ImageCard from "./imageComponents/ImageContainer";
import "../styles/Output.css"

function Output({output}){
    // const images = output["images"];
    // const uploadedFile = output["inputFile"];
    console.log("The ouptut part is: ", output);

    return (
        <>
            {output.map((data) => (
            <div className="input-file-section" key={data.inputFile.id}>

            <h2>{data.inputFile.name}</h2>

            <div className="image-grid">
                {data.images.map((image) => (
                    <ImageCard
                        key={image.id}
                        image={image}
                    />
                ))}
            </div>
        </div>
    ))}
        </>
    )
}

export default Output;