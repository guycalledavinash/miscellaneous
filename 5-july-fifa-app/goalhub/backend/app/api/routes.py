from fastapi import APIRouter, Depends, HTTPException, Query
from prometheus_client import CONTENT_TYPE_LATEST, generate_latest
from sqlalchemy.orm import Session
from starlette.responses import Response
from app.core.security import create_access_token, require_admin, verify_password
from app.db.session import get_db
from app.models.entities import Match, Player, Team, User
from app.repositories.base import Repository
from app.schemas.schemas import Dashboard, Login, MatchCreate, MatchRead, MatchUpdate, PlayerCreate, PlayerRead, PlayerUpdate, TeamCreate, TeamRead, TeamUpdate, Token
from app.services.services import DashboardService, PlayerService

router = APIRouter()
@router.get("/health")
def health(): return {"status": "ok"}
@router.get("/metrics")
def metrics(): return Response(generate_latest(), media_type=CONTENT_TYPE_LATEST)
@router.post("/auth/login", response_model=Token)
def login(payload: Login, db: Session = Depends(get_db)):
    user = db.query(User).filter(User.username == payload.username).first()
    if not user or not verify_password(payload.password, user.hashed_password): raise HTTPException(401, "Invalid credentials")
    return Token(access_token=create_access_token(user.username, user.role.value))
@router.get("/dashboard", response_model=Dashboard)
def dashboard(db: Session = Depends(get_db)): return DashboardService(db).get()

@router.get("/players", response_model=list[PlayerRead])
def players(q: str | None = None, country: str | None = None, position: str | None = None, sort: str | None = Query(None, pattern="^(goals|assists)$"), skip: int = 0, limit: int = 50, db: Session = Depends(get_db)): return PlayerService(db).search(q, country, position, sort, skip, limit)
@router.get("/players/{id}", response_model=PlayerRead)
def player(id: int, db: Session = Depends(get_db)):
    obj = Repository(db, Player).get(id)
    if not obj: raise HTTPException(404, "Player not found")
    return obj
@router.post("/players", response_model=PlayerRead, dependencies=[Depends(require_admin)])
def create_player(p: PlayerCreate, db: Session = Depends(get_db)): return Repository(db, Player).create(p.model_dump())
@router.put("/players/{id}", response_model=PlayerRead, dependencies=[Depends(require_admin)])
def update_player(id: int, p: PlayerUpdate, db: Session = Depends(get_db)):
    repo = Repository(db, Player); obj = repo.get(id)
    if not obj: raise HTTPException(404, "Player not found")
    return repo.update(obj, p.model_dump())
@router.delete("/players/{id}", dependencies=[Depends(require_admin)])
def delete_player(id: int, db: Session = Depends(get_db)):
    repo = Repository(db, Player); obj = repo.get(id)
    if not obj: raise HTTPException(404, "Player not found")
    repo.delete(obj); return {"deleted": True}

@router.get("/teams", response_model=list[TeamRead])
def teams(skip: int = 0, limit: int = 50, db: Session = Depends(get_db)): return Repository(db, Team).list(skip, limit)
@router.post("/teams", response_model=TeamRead, dependencies=[Depends(require_admin)])
def create_team(t: TeamCreate, db: Session = Depends(get_db)): return Repository(db, Team).create(t.model_dump())
@router.put("/teams/{id}", response_model=TeamRead, dependencies=[Depends(require_admin)])
def update_team(id: int, t: TeamUpdate, db: Session = Depends(get_db)):
    repo = Repository(db, Team); obj = repo.get(id)
    if not obj: raise HTTPException(404, "Team not found")
    return repo.update(obj, t.model_dump())
@router.delete("/teams/{id}", dependencies=[Depends(require_admin)])
def delete_team(id: int, db: Session = Depends(get_db)):
    repo = Repository(db, Team); obj = repo.get(id)
    if not obj: raise HTTPException(404, "Team not found")
    repo.delete(obj); return {"deleted": True}

@router.get("/matches", response_model=list[MatchRead])
def matches(skip: int = 0, limit: int = 50, db: Session = Depends(get_db)): return Repository(db, Match).list(skip, limit)
@router.post("/matches", response_model=MatchRead, dependencies=[Depends(require_admin)])
def create_match(m: MatchCreate, db: Session = Depends(get_db)): return Repository(db, Match).create(m.model_dump())
@router.put("/matches/{id}", response_model=MatchRead, dependencies=[Depends(require_admin)])
def update_match(id: int, m: MatchUpdate, db: Session = Depends(get_db)):
    repo = Repository(db, Match); obj = repo.get(id)
    if not obj: raise HTTPException(404, "Match not found")
    return repo.update(obj, m.model_dump())
@router.delete("/matches/{id}", dependencies=[Depends(require_admin)])
def delete_match(id: int, db: Session = Depends(get_db)):
    repo = Repository(db, Match); obj = repo.get(id)
    if not obj: raise HTTPException(404, "Match not found")
    repo.delete(obj); return {"deleted": True}
