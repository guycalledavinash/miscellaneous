from sqlalchemy import func, or_
from sqlalchemy.orm import Session
from app.models.entities import Match, Player, Team
from app.repositories.base import Repository

class PlayerService:
    def __init__(self, db: Session): self.repo = Repository(db, Player); self.db = db
    def search(self, q=None, country=None, position=None, sort=None, skip=0, limit=50):
        query = self.db.query(Player)
        if q: query = query.filter(or_(Player.first_name.ilike(f"%{q}%"), Player.last_name.ilike(f"%{q}%"), Player.club.ilike(f"%{q}%")))
        if country: query = query.filter(Player.nationality == country)
        if position: query = query.filter(Player.position == position)
        if sort == "goals": query = query.order_by(Player.goals.desc())
        if sort == "assists": query = query.order_by(Player.assists.desc())
        return query.offset(skip).limit(limit).all()

class DashboardService:
    def __init__(self, db: Session): self.db = db
    def get(self):
        distribution = dict(self.db.query(Player.nationality, func.count(Player.id)).group_by(Player.nationality).all())
        return {"total_players": self.db.query(Player).count(), "total_teams": self.db.query(Team).count(), "upcoming_matches": self.db.query(Match).filter(Match.home_score.is_(None)).count(), "top_scorers": self.db.query(Player).order_by(Player.goals.desc()).limit(5).all(), "top_assists": self.db.query(Player).order_by(Player.assists.desc()).limit(5).all(), "team_rankings": self.db.query(Team).order_by(Team.fifa_ranking.asc()).limit(10).all(), "nationality_distribution": distribution}
