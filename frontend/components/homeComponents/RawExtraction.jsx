import "../../styles/Home.css"

function RawExtractionCard() {
  return (
    <div className="asset-card">
      <div className="asset-card-header">
        <div className="asset-icon">☁</div>
      </div>

      <h3>RAW Extraction</h3>

      <p>
        Fetch processed raw camera buffers whose images you wanted.
      </p>

      <input type="file" className="uploadFile"/>
      <button className="uploadraw-button">
        ▧ &nbsp; Upload Local RAW (card.raw)
      </button>

      <div className="card-footer">
        <span>Upload valid .raw file to get the response.</span>
      </div>
    </div>
  )
}

export default RawExtractionCard;