from datetime import datetime
from pydantic import BaseModel, ConfigDict, Field

class Token(BaseModel):
    access_token: str
    token_type: str = "bearer"

class Login(BaseModel):
    username: str
    password: str

class PlayerBase(BaseModel):
    first_name: str
    last_name: str
    nationality: str
    club: str
    position: str
    age: int = Field(ge=15, le=50)
    goals: int = Field(ge=0)
    assists: int = Field(ge=0)
    market_value: float = Field(ge=0)
class PlayerCreate(PlayerBase): pass
class PlayerUpdate(PlayerBase): pass
class PlayerRead(PlayerBase):
    id: int
    model_config = ConfigDict(from_attributes=True)

class TeamBase(BaseModel):
    name: str
    country: str
    coach: str
    fifa_ranking: int = Field(ge=1)
class TeamCreate(TeamBase): pass
class TeamUpdate(TeamBase): pass
class TeamRead(TeamBase):
    id: int
    model_config = ConfigDict(from_attributes=True)

class MatchBase(BaseModel):
    home_team_id: int
    away_team_id: int
    stadium: str
    kickoff_time: datetime
    home_score: int | None = None
    away_score: int | None = None
    competition_id: int
class MatchCreate(MatchBase): pass
class MatchUpdate(MatchBase): pass
class MatchRead(MatchBase):
    id: int
    model_config = ConfigDict(from_attributes=True)

class Dashboard(BaseModel):
    total_players: int
    total_teams: int
    upcoming_matches: int
    top_scorers: list[PlayerRead]
    top_assists: list[PlayerRead]
    team_rankings: list[TeamRead]
    nationality_distribution: dict[str, int]
