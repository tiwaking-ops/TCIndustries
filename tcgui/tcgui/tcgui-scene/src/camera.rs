use nalgebra::{Matrix4, Point3, Vector3};

#[derive(Debug, Clone)]
pub struct Camera3D {
    pub position: Point3<f32>,
    pub target: Point3<f32>,
    pub up: Vector3<f32>,
    pub fov_y: f32,
    pub aspect: f32,
    pub z_near: f32,
    pub z_far: f32,
}

impl Camera3D {
    pub fn new(position: Point3<f32>, target: Point3<f32>, aspect: f32) -> Self {
        Self {
            position,
            target,
            up: Vector3::y(),
            fov_y: std::f32::consts::FRAC_PI_4,
            aspect,
            z_near: 0.1,
            z_far: 1000.0,
        }
    }

    pub fn view_matrix(&self) -> Matrix4<f32> {
        Matrix4::look_at_rh(&self.position, &self.target, &self.up)
    }

    pub fn projection_matrix(&self) -> Matrix4<f32> {
        nalgebra::Perspective3::new(self.aspect, self.fov_y, self.z_near, self.z_far).into()
    }

    pub fn view_projection_matrix(&self) -> Matrix4<f32> {
        self.projection_matrix() * self.view_matrix()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_camera_matrices() {
        let cam = Camera3D::new(Point3::new(0.0, 5.0, -10.0), Point3::origin(), 16.0 / 9.0);
        let vp = cam.view_projection_matrix();
        assert_eq!(vp.shape(), (4, 4));
    }
}
