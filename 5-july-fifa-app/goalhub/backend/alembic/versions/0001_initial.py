"""initial schema
Revision ID: 0001_initial
Revises:
Create Date: 2026-07-05
"""
from alembic import op
import sqlalchemy as sa
revision='0001_initial'; down_revision=None; branch_labels=None; depends_on=None
def upgrade():
    role = sa.Enum('admin','viewer', name='role'); role.create(op.get_bind(), checkfirst=True)
    op.create_table('users', sa.Column('id', sa.Integer, primary_key=True), sa.Column('username', sa.String(80), nullable=False), sa.Column('hashed_password', sa.String(255), nullable=False), sa.Column('role', role, nullable=False)); op.create_index('ix_users_username','users',['username'], unique=True)
    op.create_table('teams', sa.Column('id', sa.Integer, primary_key=True), sa.Column('name', sa.String(120), nullable=False), sa.Column('country', sa.String(80), nullable=False), sa.Column('coach', sa.String(120), nullable=False), sa.Column('fifa_ranking', sa.Integer, nullable=False)); op.create_index('ix_teams_name','teams',['name'], unique=True)
    op.create_table('players', sa.Column('id', sa.Integer, primary_key=True), sa.Column('first_name', sa.String(80), nullable=False), sa.Column('last_name', sa.String(80), nullable=False), sa.Column('nationality', sa.String(80), nullable=False), sa.Column('club', sa.String(120), nullable=False), sa.Column('position', sa.String(40), nullable=False), sa.Column('age', sa.Integer, nullable=False), sa.Column('goals', sa.Integer, nullable=False), sa.Column('assists', sa.Integer, nullable=False), sa.Column('market_value', sa.Numeric(12,2), nullable=False))
    op.create_table('competitions', sa.Column('id', sa.Integer, primary_key=True), sa.Column('name', sa.String(120), unique=True, nullable=False))
    op.create_table('matches', sa.Column('id', sa.Integer, primary_key=True), sa.Column('home_team_id', sa.Integer, sa.ForeignKey('teams.id')), sa.Column('away_team_id', sa.Integer, sa.ForeignKey('teams.id')), sa.Column('stadium', sa.String(160), nullable=False), sa.Column('kickoff_time', sa.DateTime, nullable=False), sa.Column('home_score', sa.Integer), sa.Column('away_score', sa.Integer), sa.Column('competition_id', sa.Integer, sa.ForeignKey('competitions.id')))
def downgrade():
    for t in ['matches','competitions','players','teams','users']: op.drop_table(t)
    sa.Enum(name='role').drop(op.get_bind(), checkfirst=True)
