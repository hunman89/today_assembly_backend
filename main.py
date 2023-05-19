from fastapi import FastAPI

app = FastAPI()


@app.get("/")
async def root():
    return {"message": "Hello World"}


@app.get("/assembly/seoul")
async def seoul():
    return {"message": "seoul assembly data"}
