import React, {Component} from 'react';

import Dropzone from 'react-dropzone';

class Files extends Component {
    state = {
        accepted: [],
        rejected: []
    };

    render() {
        return <div>
            <h1>Files page</h1>
            <Dropzone
                onDrop={(accepted, rejected) => { this.setState({ accepted, rejected }); }}
            >
                <p>Try dropping some files here, or click to select files to upload.</p>
            </Dropzone>
            <div>
                <h2>Accepted files</h2>
                <ul>
                    {
                        this.state.accepted.map(f => <li key={f.name}>{f.name} - {f.size} bytes</li>)
                    }
                </ul>
                <h2>Rejected files</h2>
                <ul>
                    {
                        this.state.rejected.map(f => <li key={f.name}>{f.name} - {f.size} bytes</li>)
                    }
                </ul>
            </div>
        </div>
    }
}

export default Files;