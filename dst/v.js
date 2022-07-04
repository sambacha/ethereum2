// Require Package
import postmanToOpenApi from 'postman-to-openapi'
//const postmanToOpenApi = require('postman-to-openapi')

const { promises: { readFile } } = require('fs')
const path = require('path')

// Postman Collection Path
const postmanCollection = readFileSync('./besu.json', 'utf8')

// Output OpenAPI Path
const outputFile = './collection.yml'

// Async/await
try {
    const result = await postmanToOpenApi(postmanCollection, outputFile,  { servers: [] })
    // Without save the result in a file
   // const result2 = await postmanToOpenApi(postmanCollection, null, { defaultTag: 'General' })
    console.log(`OpenAPI specs: ${result}`)
} catch (err) {
    console.log(err)
}
