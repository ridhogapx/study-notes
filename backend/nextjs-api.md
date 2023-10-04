# Next JS API Routes
Ini adalah contoh minimal / simple untuk membuat routes API di NextJS.

# Struktur Folder
<p>Seperti biasa init folder NextJS. Lalu buat folder pages/api. Berikut contohnya</p>
<img src="img/next-1.png">
<p>Berdasarkan struktur folder di atas, akan menghasilkan url berikut http://localhost:3000/api/login dan http://localhost:3000/api/v1</p>

# Handler 
Contoh handler function pada file pages/api/login.ts


```typescript
import { NextApiRequest, NextApiResponse } from "next"
import axios from "axios"

export default async function handler(req: NextApiRequest, res: NextApiResponse ) {
    if (req.method === "POST") {
        const payload = {
            email: req.body.email,
            password: req.body.password
        }
    
        try {
            const callsAPI = await axios.post("https://BACKEND_URL/api/v1/auth/signin", {
                        email: payload.email,
                        password: payload.password
            }) 

            res.json({
                data: callsAPI.data
            })
         
        } catch(err) {
            res.json({
                message: err.response.data
            })
        }

    }
}

```

# Client Consume 
Berikut adalah contoh simple consume API pada App Router.

```typescript
// Anggap saja ini adalah state data dari Form input

interface payloadAPI {
    email: string,
    password: string
}

const data = useState<payloadAPI>({
        email: "",
        password: "",
    })

// Update State di onChange form 
// ....

// Trigger API Calls 
const handleBtn = async(e: any) => {
    e.preventDefault()
    // Kita arahkan url Axios ke NextJS internal API

    try {
        const login = async axios.post("/api/login", {
           // Menggunakan data state 
            email: data.email
            password: data.password
        })
    } catch(err) {
        // handling error
    }

    
}


```

