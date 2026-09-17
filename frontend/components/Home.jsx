import "../styles/Home.css";
import Navbar from "./homeComponents/Navbar";
import RawExtractionCard from "./homeComponents/RawExtraction";
import AddImageCard from "./homeComponents/ImageuploadSection";
import { useState } from "react";


function SourceAssets() {
    // set the states of the main and subsections here
    const[main, setMain] = useState(null);
    const[subMain, setSubMain] = useState(null);


  return (
    <section className="section">
      <div className="section-heading">
        <h2>Source Assets</h2>
        <span>
          STAGE 01 / INPUT
        </span>
      </div>

      <div className="asset-grid">
        <RawExtractionCard />
        <AddImageCard main = {main}  subMain = {subMain}  setMain={setMain} setSubMain={setSubMain}/>
      </div>
    </section>
  );
}


function AttachedImageBadge() {
  return (
    <div className="attached-row">
      <div className="attached-file">
        ▧ &nbsp; image_reference.png attached (4.2 MB)
      </div>

      <span className="ready-status">
        ✓ &nbsp; READY FOR SYNTHESIS
      </span>
    </div>
  );
}


function SuggestionChips() {
  return (
    <div className="chip-row">
      <span className="chip-label">SUGGESTIONS:</span>

      <button className="chip">
        Make the image brighter and remove the background
      </button>

      <button className="chip">
        Dramatic golden hour bokeh
      </button>
    </div>
  );
}


function AspectRatioChips() {
  return (
    <div className="aspect-row">
      <span className="chip-label">ASPECT:</span>

      <button className="aspect-chip active">
        1:1 Square
      </button>

      <button className="aspect-chip">
        16:9 Cinema
      </button>

      <button className="aspect-chip">
        9:16 Story
      </button>

      <button className="aspect-chip">
        4:5 Portrait
      </button>
    </div>
  );
}


function PromptSection() {
  return (
    <section className="section prompt-section">
      <div className="section-heading">
        <h2>
          Tell the AI what you want
          <span className="heading-subtitle">
            (Image Ingestion & Prompting)
          </span>
        </h2>

        <span>
          STAGE 02 / PROMPT & CONFIG
        </span>
      </div>

      <div className="prompt-card">

        <AttachedImageBadge />

        <textarea
          className="prompt-input"
          placeholder="Describe what you want to create or change..."
        />

        <SuggestionChips />

        <div className="prompt-bottom">
          <AspectRatioChips />

          <button className="generate-button">
            ✨ &nbsp; Generate
          </button>
        </div>

      </div>
    </section>
  );
}


function RecentGeneration() {
  return (
    <section className="recent-generation">

      <div className="recent-icon">
        ▧
      </div>

      <div className="recent-info">
        <span className="live-label">
          ● LIVE BUFFER
        </span>

        <h3>
          Recent generation ready: 4096×2160 UHD
        </h3>

        <p>
          Cyberpunk architectural render • Seed #8246204
        </p>
      </div>

      <div className="recent-actions">
        <button>
          View Generated Output →
        </button>

        <button>
          Browse History Archive →
        </button>
      </div>

    </section>
  );
}


function Footer() {
  return (
    <footer className="footer">

      <span>
        © 2026 PixelForge AI Image Studio. All rights reserved.
      </span>

      <div>
        <span>v2.4.0</span>
        <span>Precision Generation Engine</span>
      </div>

    </footer>
  );
}


export default function Home() {
  return (
    <main className="home-page">

      <Navbar />

      <SourceAssets />

      <PromptSection />

      <RecentGeneration />

      <Footer />

    </main>
  );
}