from datetime import datetime, timedelta
import random
from app.core.security import hash_password
from app.db.session import Base, SessionLocal, engine
from app.models.entities import Competition, Match, Player, Role, Team, User
Base.metadata.create_all(bind=engine)
db=SessionLocal()
if not db.query(User).first():
    db.add_all([User(username='admin', hashed_password=hash_password('admin123'), role=Role.admin), User(username='viewer', hashed_password=hash_password('viewer123'), role=Role.viewer)])
competitions=[Competition(name=n) for n in ['World Cup','Champions League','Premier League','La Liga']]
if not db.query(Competition).first(): db.add_all(competitions); db.commit()
countries=['Argentina','Brazil','France','England','Spain','Germany','Portugal','Netherlands','Italy','Croatia','Japan','USA','Mexico','Morocco','Belgium','Uruguay','Colombia','Senegal','Korea Republic','Switzerland']
if not db.query(Team).first():
    teams=[Team(name=f'{c} National Team', country=c, coach=f'Coach {i+1}', fifa_ranking=i+1) for i,c in enumerate(countries)]
    db.add_all(teams); db.commit()
first=['Lionel','Kylian','Jude','Vinicius','Erling','Harry','Kevin','Luka','Pedri','Bukayo','Phil','Antoine','Bruno','Lautaro','Rodri','Son','Christian','Achraf','Federico','Rafael']
last=['Silva','Martinez','Garcia','Santos','Mbappe','Messi','Bellingham','Kane','Haaland','Modric','Foden','Saka','Griezmann','Fernandes','Valverde','Leao','Pulisic','Hakimi','Kim','Diaz']
positions=['GK','DF','MF','FW']
teams=db.query(Team).all()
if db.query(Player).count()<200:
    for i in range(200): db.add(Player(first_name=random.choice(first), last_name=f'{random.choice(last)} {i}', nationality=random.choice(countries), club=random.choice(teams).name, position=random.choice(positions), age=random.randint(18,38), goals=random.randint(0,45), assists=random.randint(0,30), market_value=random.randint(1,180)*1_000_000))
    db.commit()
comps=db.query(Competition).all()
if db.query(Match).count()<50:
    for i in range(50):
        h,a=random.sample(teams,2); played=i<15
        db.add(Match(home_team_id=h.id, away_team_id=a.id, stadium=f'Global Arena {i+1}', kickoff_time=datetime.utcnow()+timedelta(days=i-10), home_score=random.randint(0,4) if played else None, away_score=random.randint(0,4) if played else None, competition_id=random.choice(comps).id))
    db.commit()
db.close()
