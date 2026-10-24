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

/// Orthographic camera for UI / 2D overlays or ortho projection mode.
#[derive(Debug, Clone)]
pub struct OrthoCamera3D {
    pub left: f32,
    pub right: f32,
    pub bottom: f32,
    pub top: f32,
    pub z_near: f32,
    pub z_far: f32,
}

impl OrthoCamera3D {
    pub fn new(left: f32, right: f32, bottom: f32, top: f32, z_near: f32, z_far: f32) -> Self {
        Self {
            left,
            right,
            bottom,
            top,
            z_near,
            z_far,
        }
    }

    pub fn projection_matrix(&self) -> Matrix4<f32> {
        nalgebra::Orthographic3::new(
            self.left,
            self.right,
            self.bottom,
            self.top,
            self.z_near,
            self.z_far,
        )
        .into()
    }
}

#[derive(Debug, Clone, Copy)]
pub struct Transform3D {
    pub translation: Vector3<f32>,
    pub rotation: nalgebra::UnitQuaternion<f32>,
    pub scale: Vector3<f32>,
}

impl Default for Transform3D {
    fn default() -> Self {
        Self {
            translation: Vector3::zeros(),
            rotation: nalgebra::UnitQuaternion::identity(),
            scale: Vector3::repeat(1.0),
        }
    }
}

impl Transform3D {
    pub fn new(
        translation: Vector3<f32>,
        rotation: nalgebra::UnitQuaternion<f32>,
        scale: Vector3<f32>,
    ) -> Self {
        Self {
            translation,
            rotation,
            scale,
        }
    }

    pub fn matrix(&self) -> Matrix4<f32> {
        let t = Matrix4::new_translation(&self.translation);
        let r = self.rotation.to_homogeneous();
        let s = Matrix4::new_nonuniform_scaling(&self.scale);
        t * r * s
    }
}

#[derive(Debug, Clone, Copy)]
pub struct Vertex3D {
    pub position: [f32; 3],
    pub normal: [f32; 3],
    pub color: [f32; 4],
}

unsafe impl bytemuck::Pod for Vertex3D {}
unsafe impl bytemuck::Zeroable for Vertex3D {}

impl Vertex3D {
    pub fn desc<'a>() -> wgpu::VertexBufferLayout<'a> {
        use std::mem;
        wgpu::VertexBufferLayout {
            array_stride: mem::size_of::<Vertex3D>() as wgpu::BufferAddress,
            step_mode: wgpu::VertexStepMode::Vertex,
            attributes: &[
                // position (vec3)
                wgpu::VertexAttribute {
                    format: wgpu::VertexFormat::Float32x3,
                    offset: 0,
                    shader_location: 0,
                },
                // normal (vec3)
                wgpu::VertexAttribute {
                    format: wgpu::VertexFormat::Float32x3,
                    offset: 12,
                    shader_location: 1,
                },
                // color (vec4)
                wgpu::VertexAttribute {
                    format: wgpu::VertexFormat::Float32x4,
                    offset: 24,
                    shader_location: 2,
                },
            ],
        }
    }
}

/// Retro vertex-lit shading parameters (Pre-CU SWG style flat/vertex lighting).
#[derive(Debug, Clone, Copy)]
pub struct RetroLighting {
    pub light_direction: Vector3<f32>,
    pub ambient_color: [f32; 3],
    pub light_color: [f32; 3],
}

impl Default for RetroLighting {
    fn default() -> Self {
        Self {
            light_direction: Vector3::new(0.5, 1.0, -0.5).normalize(),
            ambient_color: [0.2, 0.2, 0.25], // atmospheric blue-grey ambient
            light_color: [0.9, 0.85, 0.7],   // warm sunlight
        }
    }
}

impl RetroLighting {
    /// Compute vertex-lit flat shading color for a vertex given its normal and base color.
    pub fn shade(&self, normal: Vector3<f32>, base_color: [f32; 4]) -> [f32; 4] {
        let n = normal.normalize();
        let ndotl = n.dot(&self.light_direction).max(0.0);

        let r = base_color[0] * (self.ambient_color[0] + self.light_color[0] * ndotl);
        let g = base_color[1] * (self.ambient_color[1] + self.light_color[1] * ndotl);
        let b = base_color[2] * (self.ambient_color[2] + self.light_color[2] * ndotl);

        [
            r.clamp(0.0, 1.0),
            g.clamp(0.0, 1.0),
            b.clamp(0.0, 1.0),
            base_color[3],
        ]
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

        let view = cam.view_matrix();
        assert_eq!(view.shape(), (4, 4));

        let proj = cam.projection_matrix();
        assert_eq!(proj.shape(), (4, 4));
    }

    #[test]
    fn test_ortho_camera() {
        let ortho = OrthoCamera3D::new(-1.0, 1.0, -1.0, 1.0, 0.1, 100.0);
        let proj = ortho.projection_matrix();
        assert_eq!(proj.shape(), (4, 4));
    }

    #[test]
    fn test_transform3d() {
        let transform = Transform3D::new(
            Vector3::new(1.0, 2.0, 3.0),
            nalgebra::UnitQuaternion::from_axis_angle(
                &Vector3::y_axis(),
                std::f32::consts::FRAC_PI_2,
            ),
            Vector3::repeat(2.0),
        );
        let m = transform.matrix();
        assert_eq!(m.shape(), (4, 4));
    }

    #[test]
    fn test_retro_lighting() {
        let lighting = RetroLighting::default();
        let normal = Vector3::new(0.0, 1.0, 0.0);
        let base_color = [1.0, 1.0, 1.0, 1.0];
        let shaded = lighting.shade(normal, base_color);
        assert_eq!(shaded.len(), 4);
        assert!(shaded[0] >= 0.0 && shaded[0] <= 1.0);
    }
}
