import "../../styles/ImageCard.css"

export default function ImageCard({key, image }) {
    return (
        <div className="image-card" key={key}>

            {/* Image */}
            <div className="image-container">
                <img src={image["storageKey"]} alt={image["filename"]} />
            </div>

            {/* Basic information */}
            <div className="image-details">
                <h3>{image["filename"]}</h3>
                <div className="metadata">

                    <div>
                        <span>Type</span>
                        <strong>{image["mimetype"]}</strong>
                    </div>

                    <div>
                        <span>Size</span>
                        <strong>
                            {(image["filesize"] / 1024).toFixed(2)} KB
                        </strong>
                    </div>

                    <div>
                        <span>Dimensions</span>
                        <strong>
                            {image["width"]} × {image["height"]}
                        </strong>
                    </div>
{/* 
                    <div>
                        <span>Created</span>
                        <strong>
                            {new Date(image["createdAt"]).toLocaleString()}
                        </strong>
                    </div> */}

                </div>
            </div>

            {/* Actions */}
            <div className="image-actions">
                <button>Preview</button>
                <button>Download</button>
            </div>

        </div>
    );
}