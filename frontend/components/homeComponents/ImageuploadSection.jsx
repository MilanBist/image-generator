import '../../styles/Home.css'
import MainSubMainImage from './MainSubMainImage'
// import PromptSection from './PromptSection'


export default function AddImageCard({main, subMain, setMain, setSubMain}){
    return (
        <>
            <div className="asset-card">
                <div className="asset-card-header">
                    <div className="asset-icon">☁</div>
                </div>

                <h3>Image Transformation</h3>
                <p>
                    Add .jpg, .png file to transform the image.
                </p>

                <input type="file" className="uploadImage"/>

                <MainSubMainImage main={main} subMain={subMain} setMain={setMain} setSubmain={setSubMain}/>
                <button className="uploadraw-button">
                    ▧ &nbsp; Upload image file.
                </button>

                <div className="card-footer">
                    <span>Upload valid image file.</span>
                </div>


                {/* <PromptSection/> */}
            </div>
        </>
    )
}