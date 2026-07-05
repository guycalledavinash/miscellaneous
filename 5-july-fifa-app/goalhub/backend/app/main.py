from fastapi import FastAPI, Request
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import JSONResponse
from app.api.routes import router
from app.core.config import get_settings
from app.core.logging import configure_logging

configure_logging(); settings = get_settings()
app = FastAPI(title=settings.app_name, version="1.0.0")
app.add_middleware(CORSMiddleware, allow_origins=settings.cors_origins, allow_credentials=True, allow_methods=["*"], allow_headers=["*"])
@app.exception_handler(Exception)
async def unhandled(_: Request, exc: Exception): return JSONResponse(status_code=500, content={"detail": "Internal server error"})
app.include_router(router)
